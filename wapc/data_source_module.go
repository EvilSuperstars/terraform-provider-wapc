package wapc

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/wapc/wapc-go"
)

func dataSourceWapcModule() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWapcModuleReadContext,

		Schema: map[string]*schema.Schema{
			"filename": {
				Type:     schema.TypeString,
				Required: true,
			},

			"input": {
				Type:     schema.TypeString,
				Required: true,
			},

			"operation": {
				Type:     schema.TypeString,
				Required: true,
			},

			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceWapcModuleReadContext(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	src := d.Get("filename").(string)
	code, err := ioutil.ReadFile(src)
	if err != nil {
		return diag.Errorf("error reading WebAssembly code (%s): %s", src, err)
	}

	module, err := wapc.New(func(msg string) { log.Printf("%s", msg) }, code, wapc.NoOpHostCallHandler)
	if err != nil {
		return diag.Errorf("error compiling WebAssembly module: %s", err)
	}
	defer module.Close()

	instance, err := module.Instantiate()
	if err != nil {
		return diag.Errorf("error instantiating WebAssembly module: %s", err)
	}
	defer instance.Close()

	operation := d.Get("operation").(string)
	payload := []byte(d.Get("input").(string))
	result, err := instance.Invoke(ctx, operation, payload)
	if err != nil {
		return diag.Errorf("error invoking WebAssembly operation (%s): %s", operation, err)
	}

	d.SetId(time.Now().UTC().String())
	d.Set("result", string(result))

	return nil
}
