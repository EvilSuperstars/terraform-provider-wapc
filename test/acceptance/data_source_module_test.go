package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestDataSourceWapcModule_basic(t *testing.T) {
	dataSourceName := "data.wapc_module.test"
	input := acctest.RandomWithPrefix("tf-acc-test")

	resource.ParallelTest(t, resource.TestCase{
		Providers: testProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceWapcModuleConfigBasic(input),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "result", fmt.Sprintf("Hello, %s", input)),
				),
			},
		},
	})
}

func testDataSourceWapcModuleConfigBasic(input string) string {
	return fmt.Sprintf(`
data "wapc_module" "test" {
  filename  = "testdata/hello.wasm"
  operation = "hello"
  input     = %[1]q
}
`, input)
}
