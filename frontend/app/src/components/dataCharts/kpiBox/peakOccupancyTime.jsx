import React, { useEffect, useState } from "react";
import KpiBox from "./kpiBox";


//props : data: []
const PeakOccupancyTime = (props) => {

  const [data, setData] = useState(props.data)
  const [value, setValue] = useState(0)

  const treatData = () => {
    let tempData = []
    let hourMap = {

    }

    props.data.map((obj) => {
      tempData.push({
        name: obj.roomName,
        data: obj.data
      })
    })
    console.log('Temp data', tempData)
    tempData.map((room) => {
      room.data.map((row) => {
        let hour = new Date(row.timestamp * 1000).getHours()
        if(hourMap[hour] === undefined){
          hourMap[hour] = 1
        }
        else{
          hourMap[hour]++
        }
      })
    })
    console.log('Hour map', hourMap)
    //Fråga mig inte vad som händer här, jag har ingen aning Mvh Oscar
    setValue(Object.keys(hourMap).reduce((a,b) => { return hourMap[a] > hourMap[b] ? a : b }))

  }

  const formatTime = (time) => {
    if(time < 10){
      return `0${time}:00`
    }
    return `${time}:00`
  }


  useEffect(() => {
    setData(props.data)
    console.log('Data loaded AverageOccupancyRate.jsx', data)
    if(data.length === 0){
      return
    }
    treatData()
  }, [props.data])

  return(
    <KpiBox 
      title="Peak Occupancy Hour"
      value={formatTime(value)} 
      unit=""
    />
  )
}

export default PeakOccupancyTime;