import React, { useEffect, useState } from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCardAlt from '../components/roomCardAlt/roomCardAlt'
import axios from 'axios'
import { Autocomplete, Checkbox, TextField } from '@mui/material'

function ListRooms() {

  //const API_URL = import.meta.env.API_URL
  //const API_KEY = import.meta.env.API_KEY

  const [windowWidth, setWindowWidth] = useState(window.innerWidth)
  
  const [data, setData] = useState([])
  const [roomNames, setRoomNames] = useState([])

  //Checkbox states
  const [availableCheckBox, setAvailableCheckBox] = useState(true)
  const [reservedCheckBox, setReservedCheckBox] = useState(false)
  const [occupiedCheckBox, setOccupiedCheckBox] = useState(false)

  //Search bar state
  const [searchValue, setSearchValue] = useState("")
  const [resultsCard, setResultsCard] = useState(undefined)

  const handleSearchValueChange = (value) => {
    setSearchValue(value)
    if(roomNames.includes(value)){
      console.log("Room found", value)
      setResultsCard(
        <RoomCardAlt
          key={`${value}-results`}
          RoomName={value}
          RoomHouse={"Building"}
          Avaiability={"Available"}
        />
      )
    } else {
      console.log("Room not found", value)
      setResultsCard(undefined)
    }
  }

  const client = axios.create({
    baseURL: "/api",
  })

  useEffect(() => {
    // Query the API, with axios
    const fetchData = async () => {
      client.get('/current').then((response) => { 
        setData(response.data);
        let tempRoom = [];
        response.data.forEach(obj => {
          console.log(obj.room)
          if (!tempRoom.includes(obj.room)){
            tempRoom.push(obj.room)
          }
          setRoomNames(tempRoom)
        })
      });
    }

    fetchData()

    data.forEach(obj => {
      console.log(obj)
    })
    console.log(roomNames)
  },[])

  useEffect(() => {
    function handleResize() {
      setWindowWidth(window.innerWidth)
    }
    window.addEventListener('resize', handleResize)
    return () => window.removeEventListener('resize', handleResize)
  },[])

  return (
    <>
      <div className='page-header' style={{marginTop: "12vh"}}>
        <h1>Rooms</h1>
      </div>
      
      <div style={{display: "flex", flexDirection:"row"}}>
        <div style={{marginLeft: "12vw"}}>
          <Checkbox checked={availableCheckBox} color="success" onChange={() => setAvailableCheckBox(!availableCheckBox)}/>
          <label>Available</label>

          <Checkbox  checked={reservedCheckBox} color="warning" onChange={() => setReservedCheckBox(!reservedCheckBox)}/>
          <label>Reserved</label>

          <Checkbox checked={occupiedCheckBox}  color="error" onChange={() => setOccupiedCheckBox(!occupiedCheckBox)}/>
          <label>Occupied</label>
        </div>
        <p style={{marginLeft: "15px", fontSize: "small"}}>*Reserved means that the room is booked but not in use</p>
      </div>

      <CardGrid>
        {
          data.map((room, index) => {
            if (room.status === "available" && availableCheckBox 
                || room.status === "occupied" && occupiedCheckBox 
                || room.status === "booked" && reservedCheckBox) {
              return (
                <RoomCardAlt
                  key={room.room}
                  RoomName={room.room}
                  RoomHouse={room.building}
                  Avaiability={room.status}
                />
              );
            } else {
              return null; // Return null for non-available rooms (optional)
            }
          })
        }
      </CardGrid>
      
    </>
  )
}

export default ListRooms
