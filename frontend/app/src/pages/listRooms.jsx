import CardGrid from '../components/cardGrid/cardGrid'
import MobileRoomCard from '../components/mobileRoomCard/mobileRoomCard'

function ListRooms() {
  return (
    <>
    <div className='page-header'>
      <h1>Available Rooms</h1>
    </div>
      <CardGrid>
        <MobileRoomCard RoomName="Vasa G-14" RoomHouse="Vasa" Avaiability="Available" />
        <MobileRoomCard RoomName="M1215B" RoomHouse="Maskinhuset" Avaiability="Available" />
      </CardGrid>
    </>
  )
}

export default ListRooms
