import React, { useEffect, useState } from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCardAlt from '../components/roomCardAlt/roomCardAlt'
import axios from 'axios'
import { Checkbox } from '@mui/material'

function ListRooms() {

  const [windowWidth, setWindowWidth] = useState(window.innerWidth)
  
  const [data, setData] = useState([])
  const [roomNames, setRoomNames] = useState([])

  //Checkbox states
  const [availableCheckBox, setAvailableCheckBox] = useState(true)
  const [unknownCheckBox, setUnknownCheckBox] = useState(true)
  const [occupiedCheckBox, setOccupiedCheckBox] = useState(true)



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

          <Checkbox checked={occupiedCheckBox}  color="error" onChange={() => setOccupiedCheckBox(!occupiedCheckBox)}/>
          <label>Occupied</label>

          <Checkbox  checked={unknownCheckBox} color="default" onChange={() => setUnknownCheckBox(!unknownCheckBox)}/>
          <label>Unknown</label>

        </div>
        <p style={{marginLeft: "15px", fontSize: "small"}}>*unknown means that a reliable reading hasn't been made in the room for some time</p>
      </div>

      <CardGrid>
        {
          data.map((room, index) => {
            if (room.status === "available" && availableCheckBox 
                || room.status === "reserved" && occupiedCheckBox 
                || room.status === "unknown" && unknownCheckBox) {
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
