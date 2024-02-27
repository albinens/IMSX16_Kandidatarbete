import React, {useEffect, useState} from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import MobileRoomCard from '../components/mobileRoomCard/mobileRoomCard'
import './searchRooms.css'

function SearchRooms() {

  const [searchQuery, setSearchQuery] = useState('')
  const [apiData, setApiData] = useState([]) // [roomName, house, avaiability
  const [searchResults, setSearchResults] = useState([])


  const handleSearch = (e) => {
    setSearchQuery(e.target.value)
    console.log(searchQuery)
  }

  const handleSearchButton = () => {
    let filteredData = fakeData.filter((room) => {
      return room.roomName.toLowerCase().includes(searchQuery.toLowerCase())
    })
    setSearchResults(filteredData)
    console.log(searchResults)
  }

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
    // fetch data from api
    // setApiData(data)
    // filter search results
    setSearchResults(fakeData)
  }, [])

  return (
    <>
      <div className='page-header'>
        <h1>Search</h1>
      </div>
      <div className='search-page-search-field-box'>
        <input 
          className="search-page-input-field" 
          type="text" 
          placeholder="Search for rooms" 
          onChange={(e) => handleSearch(e)}
        />
        <button 
          className='search-page-input-button'
          onClick={() => handleSearchButton()}
        >
            Search
        </button>
      </div>
      <CardGrid search>
        {
          searchResults.map((room, index) => {
            return (
              <>
                <MobileRoomCard 
                  key={index} 
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

export default SearchRooms
