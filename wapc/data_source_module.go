package wapc

import (
	"context"
	"io/ioutil"
	"os"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
	// src := d.Get("url").(string)

	tmpDir, err := ioutil.TempDir("", "gg")
	if err != nil {
		return []diag.Diagnostic{diag.FromErr(err)}
	}
	defer os.RemoveAll(tmpDir)

	// dst := path.Join(tmpDir, "gg.dat")
	// if err := gg.GetFile(dst, src); err != nil {
	// 	return err
	// }

	// bytes, err := ioutil.ReadFile(dst)
	// if err != nil {
	// 	return err
	// }

	d.SetId(time.Now().UTC().String())
	// d.Set("content", string(bytes))

	return nil
}
