package chartview

import (
	campaign "Cognitive-Backend-Challenge/api/model"
	"fmt"
	"strconv"
	"strings"
)

type pair struct {
	key string
	val int
}

type chartElement struct {
	name string
	data []string
}

func extractMonoChartData(chart []campaign.ChartResponse) ([]string, []string) {
	var labels []string
	var values []string

	for _, element := range chart {
		labels = append(labels, element.Key1)
		values = append(values, strconv.Itoa(element.Value))
	}

	return labels, values
}

func extractDualChartData(chart []campaign.ChartResponse) ([]string, []string, [][]string) {
	var categories []string
	var labels []string
	var dataset [][]string

	var categoriesMP = make(map[string]bool)
	var labelsMP = make(map[string]bool)

	var keys = make(map[string][]pair)

	for _, chartRes := range chart {
		if categoriesMP[chartRes.Key2] == false {
			categories = append(categories, chartRes.Key2)
			categoriesMP[chartRes.Key2] = true
		}

		if labelsMP[chartRes.Key1] == false {
			labels = append(labels, chartRes.Key1)
			labelsMP[chartRes.Key1] = true
		}

		keys[chartRes.Key2] = append(keys[chartRes.Key2], pair{key: chartRes.Key1, val: chartRes.Value})
	}

	for _, label := range labels {
		var tmp []string
		for _, category := range categories {
			var f bool
			for _, cate := range keys[category] {
				if label == cate.key {
					tmp = append(tmp, strconv.Itoa(cate.val))
					f = true
					break
				}
			}
			if f == false {
				tmp = append(tmp, "0")
			}
		}
		dataset = append(dataset, tmp)
	}

	return categories, labels, dataset
}

func getCategorieString(cates []string) string {
	var ret string

	for _, cate := range cates {
		ret += "'" + cate + "',"
	}

	return ret
}

// GenerateDualChart A function to generate a chart based on the given data with 2 keys
func GenerateDualChart(chart []campaign.ChartResponse) string {
	categories, labels, dataset := extractDualChartData(chart)

	var data []chartElement

	for i, element := range dataset {
		obj := chartElement{
			name: labels[i],
			data: element,
		}

		data = append(data, obj)
	}

	htmlScript1 := `<!DOCTYPE html>
	<head>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/core.js"></script>
	<script src="https://code.highcharts.com/highcharts.js"></script>
	<script src="https://code.highcharts.com/modules/exporting.js"></script>
	<script src="https://code.highcharts.com/modules/export-data.js"></script>
	</head>
	<body>
		<div id="container" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<script>
		Highcharts.chart('container', {
			chart: {
				type: 'column'
			},
			title: {
				text: 'Campaign Analysis'
			},
			subtitle: {
				text: 'Source: campaignapi.com'
			},
			xAxis: {
				categories: [
					` + getCategorieString(categories) + `
				],
				crosshair: true
			},
			yAxis: {
				min: 0,
				title: {
					text: 'Dimensions'
				}
			},
			tooltip: {
				headerFormat: '<span style="font-size:10px">{point.key}</span><table>',
				pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f} mm</b></td></tr>',
				footerFormat: '</table>',
				shared: true,
				useHTML: true
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0
				}
			},
		`

	htmlScript2 := `
		series: [
		`
	for _, row := range data {
		htmlScript2 += fmt.Sprintf(`
		{
			name: '%s',
			data: [%s] 
		},
		`, row.name, strings.Join(row.data, ","))
	}

	htmlScript3 := `
		]
		});
		</script>
	</body>
	`

	return htmlScript1 + htmlScript2 + htmlScript3
}

// GenerateMonoChart A function to generate a chart based on the given data with 1 keys
func GenerateMonoChart(chart []campaign.ChartResponse) string {
	labels, values := extractMonoChartData(chart)

	htmlScript1 := `<!DOCTYPE html>
	<head>
	<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.4.1/core.js"></script>
	<script src="https://code.highcharts.com/highcharts.js"></script>
	<script src="https://code.highcharts.com/modules/exporting.js"></script>
	<script src="https://code.highcharts.com/modules/export-data.js"></script>
	</head>
	<body>
		<div id="container" style="min-width: 310px; height: 400px; margin: 0 auto"></div>
		<script>
		Highcharts.chart('container', {
			chart: {
				type: 'column'
			},
			title: {
				text: 'Campaign Analysis'
			},
			subtitle: {
				text: 'Source: campaignapi.com'
			},
			xAxis: {
				categories: [
					` + getCategorieString(labels) + `
				],
				crosshair: true
			},
			yAxis: {
				min: 0,
				title: {
					text: 'Dimensions'
				}
			},
			tooltip: {
				headerFormat: '<span style="font-size:10px">{point.key}</span><table>',
				pointFormat: '<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f} mm</b></td></tr>',
				footerFormat: '</table>',
				shared: true,
				useHTML: true
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0
				}
			},
			`
	htmlScript2 := `
		series: [
		`

	htmlScript2 += fmt.Sprintf(`
		{
			name: 'Count',
			data: [%s] 
		},
		`, strings.Join(values, ","))

	htmlScript3 := `
		]
		});
		</script>
	</body>
	`

	return htmlScript1 + htmlScript2 + htmlScript3
}
