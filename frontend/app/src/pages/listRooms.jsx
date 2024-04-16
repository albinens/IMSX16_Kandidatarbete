import React, { useEffect, useState } from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCard from '../components/roomCard/roomCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'
import axios from 'axios'
import { Autocomplete, TextField } from '@mui/material'

function ListRooms() {

  //const API_URL = import.meta.env.API_URL
  //const API_KEY = import.meta.env.API_KEY

  const [windowWidth, setWindowWidth] = useState(window.innerWidth)
  
  const [data, setData] = useState([])
  const [roomNames, setRoomNames] = useState([])

  const client = axios.create({
    baseURL: "http://localhost:8080/api",
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
    <div className='page-header' style={windowWidth < 768 ? {marginTop:"3vh"} : {marginTop:"6vh"}}>
      <h1>Available Rooms</h1>
    </div>
    <Autocomplete
      disablePortal
      id="find-a-room-box-demo"
      options={roomNames}
      sx={{ width: 300 }}
      renderInput={(params) => <TextField {...params} label="Find a room" />}
      style={{margin: "0 auto", display: "block", width: "50%"}}
    />
    <HorizontalLegend />
      <CardGrid>
        {
          data.map((room, index) => {
            if (room.status === "available") {
              return (
                <RoomCard
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
