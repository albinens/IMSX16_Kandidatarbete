import React, { useEffect, useState } from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCard from '../components/roomCard/roomCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'
import axios from 'axios'

function ListRooms() {

  const API_URL = import.meta.env.API_URL
  const API_KEY = import.meta.env.API_KEY

  const [filteredQuery, setFilteredQuery] = useState([])
  const [windowWidth, setWindowWidth] = useState(window.innerWidth)
  
  const [data, setData] = useState([])

  const fakeData = [
    {
      roomName: 'Vasa G-14',
      house: 'Vasa',
      avaiability: 'Available'
    },
    {
      roomName: 'EG3503',
      house: 'EDIT',
      avaiability: 'Booked'
    },
    {
      roomName: 'F4058',
      house: 'Fysikhuset',
      avaiability: 'Occupied'
    },
    {
      roomName: 'M1215B',
      house: 'Maskinhuset',
      avaiability: 'Available'
    },
    {
      roomName: 'SB-G303',
      house: 'SB-huset',
      avaiability: 'Available'
    },
    {
      roomName: 'M1214E',
      house: 'Maskinhuset',
      avaiability: 'Occupied'
    }
  ]

  useEffect(() => {

    // Query the API, with axios
    axios.get(`${API_URL}/rooms`).then((response) => {
      setData(response.data)
    })
    console.log("DATA:" + data)

    let filteredData = fakeData.filter((room) => {
      return room.avaiability.toLowerCase().includes("available")
    })
    setFilteredQuery(filteredData)

    // Window resize event listener
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
    <HorizontalLegend />
      <CardGrid>
        {
          filteredQuery.map((room, index) => {
            return (
              <>
                <RoomCard 
                  key={room.roomName} 
                  RoomName={room.roomName} 
                  RoomHouse={room.house} 
                  Avaiability={room.avaiability} 
                />
              </>
            )
          })
        }
      </CardGrid>
    </>
  )
}

export default ListRooms
