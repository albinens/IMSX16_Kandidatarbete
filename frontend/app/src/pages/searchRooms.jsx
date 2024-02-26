import CardGrid from '../components/cardGrid/cardGrid'
import MobileRoomCard from '../components/mobileRoomCard/mobileRoomCard'

function SearchRooms() {
  return (
    <>
      <h1>Search Rooms</h1>
      <div className='search-field-box'>
        <input type="text" placeholder="Search for rooms" />
        <button>Search</button>
      </div>
      <h1 style={{textAlign:"center"}}>Results</h1>
      <CardGrid search>
        <MobileRoomCard RoomName="Vasa G-14" RoomHouse="Vasa" Avaiability="Available" />
        <MobileRoomCard RoomName="EG3503" RoomHouse="EDIT" Avaiability="Booked" />
        <MobileRoomCard RoomName="M1215A" RoomHouse="Maskinhuset" Avaiability="Occupied" />
        <MobileRoomCard RoomName="M1215B" RoomHouse="Maskinhuset" Avaiability="Available" />
      </CardGrid>
    </>
  )
}

export default SearchRooms
