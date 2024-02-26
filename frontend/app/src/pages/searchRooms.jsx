import CardGrid from '../components/cardGrid/cardGrid'
import MobileRoomCard from '../components/mobileRoomCard/mobileRoomCard'

function SearchRooms() {
  return (
    <>
      <CardGrid>
        <MobileRoomCard RoomName="Vasa G-14" RoomHouse="Vasa" Avaiability="Available" />
        <MobileRoomCard RoomName="EG3503" RoomHouse="EDIT" Avaiability="Booked" />
        <MobileRoomCard RoomName="M1215A" RoomHouse="Maskinhuset" Avaiability="Occupied" />
        <MobileRoomCard RoomName="M1215B" RoomHouse="Maskinhuset" Avaiability="Available" />
      </CardGrid>
    </>
  )
}

export default SearchRooms
