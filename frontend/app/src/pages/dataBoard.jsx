import React, { useEffect, useMemo } from 'react';
import axios from 'axios'
import './styles/dataBoard.css';

import ChartContainer from '../components/dataCharts/chartContainer/chartContainer'
import AverageOccupancyRate from '../components/dataCharts/kpiBox/averageOccupancyRate';
import PeakOccupancyTime from '../components/dataCharts/kpiBox/peakOccupancyTime';
import AverageOccupancyDuration from '../components/dataCharts/kpiBox/averageOccupancyDuration';
import useAuth from '../hooks/useAuth';

const DataBoard = () => {

  const [isLoading, setIsLoading] = React.useState(true)
  const [datasetTable, setDatasetTable] = React.useState(undefined)

  const { authenticated, authKey, handleAuthKeySet } = useAuth()

  const unixTimeNow = Math.floor(Date.now() / 1000)
  const unixTimeWeekAgo = unixTimeNow - 604800

  const client = axios.create({
    baseURL: "/api",
    headers: {
      'X-API-KEY': authKey
    }
  })

  //Graph 1 data load (right column)
  const [datasetGraph, setDatasetGraph] = React.useState([])
  useMemo(() => {
    const fetchDataGraph = async () => {
      let route = `/stats/raw-serial/${unixTimeWeekAgo}/${unixTimeNow}/1m`
      client.get(route)
      .then((response) => {
        setDatasetGraph(response.data)
        setIsLoading(false)
      }).catch((err) => {
        console.log(`Error fetching data ${err}`)
      })
    }
    if(authenticated){
      fetchDataGraph()
    }
  },[isLoading, authenticated])

  //Graph 2 data load (right column)
  const [datasetGraph2, setDatasetGraph2] = React.useState([])
  useMemo(() => {
    const fetchDataGraph2 = async () => {
      let route = `/stats/raw-serial/${unixTimeWeekAgo}/${unixTimeNow}/1d`
      client.get(route)
      .then((response) => {
        setDatasetGraph2(response.data)
        setIsLoading(false)
      }).catch((err) => {
        console.log(`Error fetching data ${err}`)
      })
    }
    if(authenticated){
      fetchDataGraph2()
    }
  },[isLoading, authenticated])

  //Table data load (left column)
  useEffect(() => {
    setIsLoading(true)
    const fetchDataTable = async () => {
      client.get(`/stats/daily-average/${unixTimeWeekAgo}/${unixTimeNow}`)
      .then((response) => {
        let data = response.data
        //Sort based on name
        data.sort((a, b) => (a.roomName > b.roomName) ? 1 : -1)
        setDatasetTable(data)

      }).catch((err) => {
        console.log(`Error fetching data ${err}`)
      })
    }
    if(authenticated){
      fetchDataTable()
    }
  }, [authenticated])


  return (
    <>
    {
      //Very simple authentication to protect everything
      //Data resolution
      !authenticated ? 
        <div className='page-header'> 
          <h2>Not Authenticated</h2> 
          <input type='password' placeholder='Enter password' onChange={(e) => {
            handleAuthKeySet(e.target.value)
          }} 
          />
        </div> : 
      <>
      <div className='page-header'>
        <h1>Data Insights</h1>
      </div>
      <div className='two-column-wrapper-dataBoard'>
        {/* LEFT COLUMN */ }
        <div className='left-column-dataBoard'>
          <div style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'space-between',
            marginBottom: '20px'
          }}>
            <AverageOccupancyRate title='Total Occupancy' data={datasetGraph2} />
            <PeakOccupancyTime title='Peak Occupancy Time' data={datasetGraph} />
            <AverageOccupancyDuration title='Average Occupancy Duration' data={datasetGraph} />
          </div>
          <table className='week-table-dataBoard'>
            <thead>
              <tr>
                <th>Room</th>
                <th>Monday</th>
                <th>Tuesday</th>
                <th>Wednesday</th>
                <th>Thursday</th>
                <th>Friday</th>
              </tr>
            </thead>

            <tbody>
              {
                datasetTable && datasetTable.map((obj) => {
                  return (
                    <tr key={obj.roomName}>
                      <td>{obj.roomName}</td>
                      <td>{obj.dailyAverages.Monday}</td>
                      <td>{obj.dailyAverages.Tuesday}</td>
                      <td>{obj.dailyAverages.Wednesday}</td>
                      <td>{obj.dailyAverages.Thursday}</td>
                      <td>{obj.dailyAverages.Friday}</td>
                    </tr>
                  )
                })
              }
            </tbody>
          </table>
        </div>
        {/* RIGHT COLUMN */ }
        <div className='right-column-dataBoard'>
          <h2>Graphs</h2>
          <ChartContainer 
            dataSeries={datasetGraph} 
            chartHeader={"1 week overlook by hour"}
            chartID={"chart1"}
          />
          <ChartContainer 
            dataSeries={datasetGraph2} 
            chartHeader={"1 week overlook by daily average"}
            chartID={"chart2"}
          />
        </div>
      </div>
      </>
    }
    </>
  );
}

export default DataBoard;