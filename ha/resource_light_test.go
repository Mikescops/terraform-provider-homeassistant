package ha

import (
	"fmt"
	"testing"

	hac "terraform-provider-ha/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccHaLightBasic(t *testing.T) {
	lightID := "light.dummy_light"
	state := "on"
	rgbColor := []int{255, 0, 0}
	brightness := 100

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckHaLightDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckHaLightConfigBasic(lightID, state, rgbColor, brightness),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckHaLightExists("ha_light.new"),
				),
			},
		},
	})
}

func testAccCheckHaLightDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*hac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ha_light" {
			continue
		}

		lightID := rs.Primary.ID

		_, err := c.SetLightState(hac.LightParams{ID: lightID}, "off")
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckHaLightConfigBasic(lightID string, state string, rgbColor []int, brightness int) string {
	return fmt.Sprintf(`
	resource "ha_light" "new" {
		entity_id = "%s"
		state     = "%s"

		rgb_color  = [%d, %d, %d]
		brightness = %d
	}
	`, lightID, state, rgbColor[0], rgbColor[1], rgbColor[2], brightness)
}

func testAccCheckHaLightExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No LightID set")
		}

		return nil
	}
}
