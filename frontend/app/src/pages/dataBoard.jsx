import React, { useEffect, useMemo } from 'react';
import axios from 'axios'
import './styles/dataBoard.css';

import ChartContainer from '../components/dataCharts/chartContainer/chartContainer'

const DataBoard = () => {

  const [isLoading, setIsLoading] = React.useState(true)
  const [datasetGraph, setDatasetGraph] = React.useState([])
  const [datasetTable, setDatasetTable] = React.useState(undefined)

  const [authenticated, setAuthenticated] = React.useState(false)
  const authCode = 'super_secret_key'

  const client = axios.create({
    baseURL: "http://localhost:8080/api",
    headers: {
      'X-API-KEY': 'super_secret_key'
    }
  })

  //Graph data load (right column)
  useMemo(() => {
    const fetchDataGraph = async () => {
      let route = `/stats/raw-serial/1712860378/1713860378/5m`
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
    console.log('Data loaded dataBoard.jsx', datasetGraph)
  },[isLoading, authenticated])

  //Table data load (left column)
  useEffect(() => {
    setIsLoading(true)
    const fetchDataTable = async () => {
      client.get('/stats/daily-average/1614556800000/1614643200000')
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
      !authenticated ? 
        <div className='page-header'> 
          <h2>Not Authenticated</h2> 
          <input type='password' placeholder='Enter password' onChange={(e) => {
            if(e.target.value === authCode){
              setAuthenticated(true)
            }
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
          <ChartContainer dataSeries={datasetGraph}/>
        </div>
      </div>
      </>
    }
    </>
  );
}

export default DataBoard;