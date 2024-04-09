import React, {useEffect, useState} from 'react'
import axios from 'axios'
import CardGrid from '../components/cardGrid/cardGrid'
import SensorCard from '../components/sensorCard/sensorCard'
import './styles/sensors.css'

function Sensors() {

  const [sensorAlreadyRegistered, setSensorAlreadyRegistered] = useState(false)
  const [sensorName, setSensorName] = useState("")
  const [sensorRoom, setSensorRoom] = useState("")
  const [sensorHouse, setSensorHouse] = useState("")
  const [sensorMacAddress, setSensorMacAddress] = useState("")

  const [sensorData, setSensorData] = useState([])
  const [recordedSensorNames, setRecordedSensorNames] = useState([])
  const client = axios.create({
    baseURL: "http://localhost:8080/api",
  })


  useEffect(() => {
    // Query the API, with axios
    const fetchData = async () => {
      client.get('/current').then((response) => { 
        setSensorData(response.data);
      });
    }
    fetchData()
    sensorData.forEach(obj => {
      recordedSensorNames.push(obj.room)
    })
  }, [])


  const handleSubmit = (e) => {
    e.preventDefault()
    setSensorAlreadyRegistered(false)
    console.log('Sensor Name:', sensorName)
    console.log('Sensor House:', sensorHouse)
    console.log('Sensor Mac Address:', sensorMacAddress)

    if(recordedSensorNames.includes(sensorRoom)) {
      setSensorAlreadyRegistered(true)
      return;
    }
    client.post('/add-room', {
      Name: sensorName,
      Sensor: sensorMacAddress,
      Buidling: sensorHouse,
    }).then((response) => {
      console.log(response)
    })
  }

  return (
    <div>
      <div className='page-header'>
        <h1>Sensors</h1>
      </div>
      <div className='two-column-wrapper-sensors'>
        <div className='left-column-sensors'>
          <h1>List of Sensors</h1>
        <CardGrid>
          {
            sensorData.map((sensor) => {
              return (
                <SensorCard
                  key={sensor.room}
                  RoomName={sensor.room}
                  RoomHouse={sensor.building}
                  Status={"online"}
                />
              )
            })
          }
        </CardGrid>
        </div>
        <div className='right-column-sensors'>
        <h1>Register Sensor</h1>
          <form className='sensor-register-form'>
        <label className='sensor-form-lab'>
          Sensor Name
          <input
            type="text"
            className="sensor-form-input"
            value={sensorName}
            onChange={(e) => setSensorName(e.target.value)}
            placeholder='Eg. F4015-1'
          />
        </label>
        <label>
          Sensor Room
          <input
            type="text"
            className="sensor-form-input"
            value={sensorRoom}
            onChange={(e) => setSensorRoom(e.target.value)}
            placeholder='Eg. F4015'
          />
        </label>        
        <label>
          House
          <input
            type="text"
            className="sensor-form-input"
            value={sensorHouse}
            onChange={(e) => setSensorHouse(e.target.value)}
            placeholder='Fysikhuset'
          />
        </label>
        <label>
          Sensor Mac Address
          <input
            type="text"
            className="sensor-form-input"
            value={sensorMacAddress}
            onChange={(e) => setSensorMacAddress(e.target.value)}
            placeholder='XX:XX:XX:XX:XX:XX'
          />
        </label>
        <button type="submit" className="submit-button" onClick={(e) => handleSubmit(e)}>Submit</button>
          </form>
        </div>
      </div>
    </div>
  )
}

export default Sensors