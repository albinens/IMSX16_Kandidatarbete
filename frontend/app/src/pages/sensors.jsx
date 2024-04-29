import React, {useEffect, useState} from 'react'
import axios from 'axios'
import CardGrid from '../components/cardGrid/cardGrid'
import SensorCard from '../components/sensorCard/sensorCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'
import './styles/sensors.css'

function Sensors() {


  const API_KEY = import.meta.env.API_KEY

  const [authenticated, setAuthenticated] = useState(localStorage.getItem("authed") === 'true' ? true : false)
  const authCode = 'super_secret_key'

  const [sensorAlreadyRegistered, setSensorAlreadyRegistered] = useState(false)
  const [sensorName, setSensorName] = useState("")
  const [sensorRoom, setSensorRoom] = useState("")
  const [sensorHouse, setSensorHouse] = useState("")
  const [sensorMacAddress, setSensorMacAddress] = useState("")

  const [sensorData, setSensorData] = useState([])
  const [recordedSensorNames, setRecordedSensorNames] = useState([])
  const client = axios.create({
    baseURL: "/api",
    headers: {
      'X-API-KEY': 'super_secret_key'
    }
  })



  useEffect(() => {
    // Query the API, with axios
    const fetchData = async () => {
      client.get('/current').then((response) => { 
        setSensorData(response.data);
      });
    }
    if(authenticated){
      fetchData()
      sensorData.forEach(obj => {
        recordedSensorNames.push(obj.room)
      })
    }
  }, [authenticated])


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
    console.log(API_KEY)
    client.post('/add-room', {
      "name": sensorName,
      "mac-address": sensorMacAddress,
      "building": sensorHouse,
    }).then((response) => {
      console.log(response)
    }).catch((error) => {
      console.log(error)
    });
  }

  return (
    <>
    {
      !authenticated ? 
      <div className='page-header'> 
      <h2>Not Authenticated</h2> 
      <input type='password' placeholder='Enter password' onChange={(e) => {
        if(e.target.value === authCode){
          localStorage.setItem("auth", authCode);
          localStorage.setItem("authTime", Date.now());
          localStorage.setItem("authed", true)
          setAuthenticated(true)
        }
      }} 
      />
    </div> : 
    <div>
    <div className='page-header'>
      <h1>Sensors</h1>
    </div>
    <div className='two-column-wrapper-sensors'>

      {/* LEFT COLUMN */}
      <div className='left-column-sensors'>
        <h1>List of Sensors</h1>
        <HorizontalLegend green="Active" yellow="Warning" red="Not Responding"/>
      <CardGrid>
        {
          sensorData.map((sensor) => {
            return (
              <SensorCard
                key={sensor.room}
                RoomName={sensor.room}
                RoomHouse={sensor.building}
                Status={"online"}
                sensorData={sensorData}
                sensorDataSetter={setSensorData}
              />
            )
          })
        }
      </CardGrid>
      </div>

      {/* RIGHT COLUMN */}
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
    }
    </>

  )
}

export default Sensors