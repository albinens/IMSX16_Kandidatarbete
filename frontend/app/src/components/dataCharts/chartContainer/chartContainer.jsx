import React, {useState} from "react";
import WeekChart from "../weekChart/weekChart";
import Legend from "../weekChart/legend";

//SAMPLE DATA
import portfolio from "./portfolio.json";
import schc from "./schc.json";
import vcit from "./vcit.json";

const ChartContainer = ({ dataSerie }) => {

  const dimensions = {
    width: 600,
    height: 300,
    margin: {
      top: 30,
      right: 30,
      bottom: 30,
      left: 60
    }
  };

  //SAMPLE DATA
  const portfolioData = {
    name: "Portfolio",
    color: "#000000",
    items: portfolio.map((d) => ({ ...d, date: new Date(d.date) }))
  };
  const schcData = {
    name: "SCHC",
    color: "#d53e4f",
    items: schc.map((d) => ({ ...d, date: new Date(d.date) }))
  };
  const vcitData = {
    name: "VCIT",
    color: "#5e4fa2",
    items: vcit.map((d) => ({ ...d, date: new Date(d.date) }))
  };


  const [selectedItems, setSelectedItems] = useState([]);
  const legendData = [portfolioData, schcData, vcitData];
  const chartData = [
    portfolioData,
    ...[schcData, vcitData].filter((d) => selectedItems.includes(d.name))
  ];
  const onChangeSelection = (name) => {
    const newSelectedItems = selectedItems.includes(name)
      ? selectedItems.filter((item) => item !== name)
      : [...selectedItems, name];
    setSelectedItems(newSelectedItems);
  };

  return (
    <div className="chart-container">
      <Legend 
        data={legendData} 
        onChange={onChangeSelection} 
        legendData={selectedItems}
      />
      <WeekChart data={chartData} dimensions={dimensions} />
    </div>
  );
}

export default ChartContainer;