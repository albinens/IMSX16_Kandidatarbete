import React, {useEffect, useMemo, useState} from "react";
import { Chart } from 'react-google-charts';
import Legend from "../weekChart/legend";


/**
 * dataSeries format:
 * [
 *  {
 *    name: string,
 *    color: string,
 *  items: [
 *  {
 *    data: number,
 *    date: Date
 *  }
 *  ...]
 * ...]
 */
const ChartContainer = (props) => {

  const chartOptions = {
    title: "Recorded People in a Room",
    hAxis: { title: "Time", viewWindow: { min: 1712860378, max: 1713860378 } },
    vAxis: { title: "Recorded People", viewWindow: { min: 0, max: 15 } },
    legend: "none"
  };


  const [dataSeries, setDataSeries] = useState(props.dataSeries);
  const [chartData, setChartData] = useState([]);
  const [noChart, setNoChart] = useState(true);
  const [selectedItems, setSelectedItems] = useState([]);


  useEffect(() => {
    let tempDataSeries = [];
    let counter = 0;
    const colors = 
    [
    "#000000", "#d53e4f", "#5e4fa2", "#3288bd", "#66c2a5", "#abdda4", "#e6f598", "#fee08b", "#fdae61", "#f46d43",
    "#d53e4f", "#9e0142", "#5e4fa2", "#3288bd", "#66c2a5", "#abdda4", "#e6f598", "#fee08b", "#fdae61", "#f46d43",
    "#5e4fa2", "#3288bd", "#66c2a5", "#abdda4", "#e6f598", "#fee08b", "#fdae61", "#f46d43", "#d53e4f", "#9e0142",
    "#66c2a5", "#abdda4", "#e6f598", "#fee08b", "#fdae61", "#f46d43", "#d53e4f", "#9e0142", "#5e4fa2", "#3288bd"
   ];
    props.dataSeries.map((obj) => {
        tempDataSeries.push({
          name: obj.roomName,
          color: colors[counter],
          items: obj.data
        })
        counter++;
    })
    setDataSeries(tempDataSeries);
    console.log('Data series', dataSeries)

    let filteredData = dataSeries.filter((room) => selectedItems.includes(room.name))
    .map((room) => {
      return room.items.map((row) => {
        return [row.timestamp, row.occupancy]
      })
    })

    let destructedArray = []
    filteredData.map(array => destructedArray.push(...array))

    setChartData([
      ["Time", "Number of People"],
      [0, 0],
      ...destructedArray
    ])
    console.log('Chart data', chartData)

    setNoChart(chartData.length === 0);
  }, [selectedItems]);

  const legendData = dataSeries;

  const onChangeSelection = (name) => {
    const newSelectedItems = selectedItems.includes(name)
      ? selectedItems.filter((item) => item !== name)
      : [...selectedItems, name];
    setSelectedItems(newSelectedItems);
  };

  return (
    <div className="chart-container">

      {
        noChart ? <h1>No Data to Display</h1> :
        <>
          <Legend 
          data={legendData} 
          onChange={onChangeSelection} 
          selectedItems={selectedItems}
        />
          <Chart
            width={'700px'}
            height={'400px'}
            chartType="LineChart"
            data={chartData}
            options={chartOptions}
            graphID="LineChart"
          />
        </>
      }

      <div className="chart">
      </div>
    </div>
  );
}

export default ChartContainer;