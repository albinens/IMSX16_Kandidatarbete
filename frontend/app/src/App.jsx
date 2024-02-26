import { useState } from 'react'
import './App.css'

import CardGrid from './components/cardGrid/cardGrid'
import MobileRoomCard from './components/mobileRoomCard/mobileRoomCard'

function App() {

  return (
    <>
      <CardGrid>
        <MobileRoomCard RoomName="Room 1" RoomHouse="House 1" Avaiability="Available" />
        <MobileRoomCard RoomName="Room 2" RoomHouse="House 2" Avaiability="Booked" />
        <MobileRoomCard RoomName="Room 3" RoomHouse="House 3" Avaiability="Occupied" />
      </CardGrid>
    </>
  )
}

export default App
