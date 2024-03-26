import React, {useEffect, useState} from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import SensorCard from '../components/sensorCard/sensorCard'
import './styles/sensors.css'

function Sensors() {

  const [sensorAlreadyRegistered, setSensorAlreadyRegistered] = useState(false)
  const [sensorName, setSensorName] = useState("")
  const [sensorRoom, setSensorRoom] = useState("")
  const [sensorRegisteredDate, setSensorRegisteredDate] = useState("")
  const [sensorLastUpdated, setSensorLastUpdated] = useState("")
  const [sensorStatus, setSensorStatus] = useState("")
  const [sensorMacAddress, setSensorMacAddress] = useState("")

  const [sensorData, setSensorData] = useState([])

  useEffect(() => {
    //fetch sensor data from backend
    setSensorData(fakeSensorData)
  }, [])


  const handleSubmit = (e) => {
    e.preventDefault()
    //setSensorAlreadyRegistered(true), backend call to check (by MAC address) if sensor is already registered
    setSensorRegisteredDate(new Date()) //set to current date
    setSensorLastUpdated(new Date()) //set to current date
    console.log('Sensor Name:', sensorName)
    console.log('Sensor Room:', sensorRoom)
    console.log('Sensor Mac Address:', sensorMacAddress)
    console.log('Sensor Registered Date:', sensorRegisteredDate)
    console.log('Sensor Last Updated:', sensorLastUpdated)
    //Assuming everything is correct, send data to backend and await confirmation to setSensorStatus('Registered')
    //setSensorStatus('Registered')
  }

  const fakeSensorData = [
    {
      sensorName: 'F4015-1',
      sensorRoom: 'F4015',
      sensorRegisteredDate: '2021-06-01 12:00:00',
      sensorLastUpdated: '2021-06-01'
    },
    {
      sensorName: 'F4015-2',
      sensorRoom: 'F4015',
      sensorRegisteredDate: '2021-06-01 12:00:00',
      sensorLastUpdated: '2021-06-01'
    }]

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
            fakeSensorData.map((sensor) => {
              return (
                <SensorCard
                  key={sensor.sensorName}
                  RoomName={sensor.sensorName}
                  RoomHouse={sensor.sensorRoom}
                  Avaiability={"Online"}
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