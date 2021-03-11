package Employee

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"

	"strconv"

	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEmployee() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEmployeeRead,
		Schema: map[string]*schema.Schema{
			"user_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mob_no": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"position": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},

	}
}

func dataSourceEmployeeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	r, err := http.Get( "http://localhost:8080/")
	if err != nil {
		return diag.FromErr(err)
	}

	//r, err := client.Do(req)
	//if err != nil {
	//	return diag.FromErr(err)
	//}
	log.Printf("Response:%v",r)
	defer r.Body.Close()


	employee := make([]map[string]interface{}, 0)

	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		return diag.FromErr(err)
	}
	log.Printf("Response:%v",employee)

	if err := d.Set("position", employee[0]["Position"]); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}




