import React, { useEffect } from 'react';
import axios from 'axios'
import './styles/dataBoard.css';

import ChartContainer from '../components/dataCharts/chartContainer/chartContainer'

const DataBoard = () => {

  const [datasetGraph, setDatasetGraph] = React.useState(undefined)
  
  const [datasetTable, setDatasetTable] = React.useState(undefined)
  const [processedTableData, setProcessedTableData] = React.useState(undefined)

  const client = axios.create({
    baseURL: "http://localhost:8080/api",
    headers: {
      'X-API-KEY': 'super_secret_key'
    }
  })
  useEffect(() => {
    
    const fetchDataGraph = async () => {
      client.get('/stats/raw-serial/1614556800000/1614643200000/5m')
      .then((response) => {
        setDatasetGraph(response.data)
      }).catch((err) => {
        console.log(`Error fetching data ${err}`)
      })
    }

    const fetchDataTable = async () => {
      client.get('/stats/daily-average/1614556800000/1614643200000')
      .then((response) => {
        console.log(response)
        let data = response.data
        //Sort based on name
        data.sort((a, b) => (a.roomName > b.roomName) ? 1 : -1)
        setDatasetTable(data)

      }).catch((err) => {
        console.log(`Error fetching data ${err}`)
      })
    }

    fetchDataGraph()
    fetchDataTable()
  },[])


  return (
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
          <ChartContainer />
        </div>
      </div>
    </>

  );
}

export default DataBoard;