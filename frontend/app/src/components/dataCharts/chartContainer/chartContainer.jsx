import React, {useEffect, useState} from "react";
import { Chart } from 'react-google-charts';
import Legend from "../legend/legend";


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

  const [dataSeries, setDataSeries] = useState(props.dataSeries);
  const [chartData, setChartData] = useState([]);
  const [noChart, setNoChart] = useState(true);
  const [selectedItems, setSelectedItems] = useState([]);
  const [dateTicks, setDateTicks] = useState([])


  const proccessGraphData = () => {
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

    let lowestDate = 999999999;
    let filteredData = dataSeries.filter((room) => selectedItems.includes(room.name))
    .map((room) => {
      return room.items.map((row) => {
        let date = new Date(row.timestamp* 1000)
        if(row.timestamp < lowestDate) lowestDate = row.timestamp
        return [date, row.occupancy, room.color]
      })
    })
    for(let i = 0; i < 2; i++){
      setDateTicks([ ...dateTicks, dateTicks.push(new Date((lowestDate + (i * 86400)) * 1000))])
    }
    let destructedArray = []
    filteredData.map(array => destructedArray.push(...array))

    setChartData([
      ["Time", "Number of People", { role: "style" }],
      ...destructedArray || [0,0,0]
    ])
  }

  useEffect(() => {
    proccessGraphData();
    setNoChart(chartData.length === 0);
  }, [selectedItems, props.dataSeries]);

  const legendData = dataSeries;

  const onChangeSelection = (name) => {
    const newSelectedItems = selectedItems.includes(name)
      ? selectedItems.filter((item) => item !== name)
      : [...selectedItems, name];
    setSelectedItems(newSelectedItems);
  };

  let chartOptions = {
    title: props.chartHeader,
    hAxis: { title: "Time", format: "dd-MM-yyyy HH:mm"},
    vAxis: { title: "Recorded People", viewWindow: { min: 0, max: 10 } },
    legend: "none",
    ticks: dateTicks
  };

  return (
    <div className="chart-container" style={{width: "90%"}}>

      {
        noChart ? <h1>No Data to Display</h1> :
        <>
          <Legend 
            data={legendData} 
            onChange={onChangeSelection} 
            selectedItems={selectedItems}
            style={{left: "0px", top: "0px"}}
          />
            <div style={{height: "400px"}}>
            {
              selectedItems.length === 0 && 
              <h2>Select a room to display</h2> ||
              <Chart
              width={'100%'}
              height={'100%'}
              chartType="ColumnChart"
              data={chartData}
              options={chartOptions}
              graphID={props.chartID}
            />
            }
            </div>

        </>
      }
    </div>
  );
}

export default ChartContainer;