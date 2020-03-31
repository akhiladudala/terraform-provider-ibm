package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMIAMAccessGroupDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMIAMAccessGroupDataSourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_iam_access_group.accgroupdata", "access_group_name", name),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupDataSourceConfig(name string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group" "accgroup" {
		name = "%s"
		tags = ["tag1", "tag2"]
	  }
	  
	  
	  data "ibm_iam_access_group" "accgroupdata" {
		access_group_name = "${ibm_iam_access_group.accgroup.name}"
	  }
	  `, name)

}