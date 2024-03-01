import React, { useEffect, useState } from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCard from '../components/roomCard/roomCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'

function ListRooms() {

  const [filteredQuery, setFilteredQuery] = useState([])
  const [windowWidth, setWindowWidth] = useState(window.innerWidth)

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
