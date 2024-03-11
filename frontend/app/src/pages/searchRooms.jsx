import React, {useEffect, useState} from 'react'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCard from '../components/roomCard/roomCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'
import './styles/search.css'

function SearchRooms() {

  const [searchQuery, setSearchQuery] = useState('')
  const [apiData, setApiData] = useState([]) // [roomName, house, avaiability
  const [searchResults, setSearchResults] = useState([])


  const handleSearch = (e) => {
      setSearchQuery(e.target.value)
      console.log(searchQuery)
      if(e.target.value === ''){
        setSearchResults(fakeData) //Change to API data
      }
  }

  const handleSearchButton = () => {
    let filteredData = fakeData.filter((room) => { //Change to API data
      return room.roomName.toLowerCase().includes(searchQuery.toLowerCase())
    })
    setSearchResults(filteredData)
    console.log(searchResults)
  }

  const handleEnterSearch = (e) => {
    if (e.key === 'Enter') {
      handleSearchButton()
    }
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
    setSearchResults(fakeData) //Change to API data
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
          onKeyDown={(e) => handleEnterSearch(e)}
        />
        <button 
          className='search-page-input-button'
          onClick={() => handleSearchButton()}
        >
            Search
        </button>
      </div>
      <HorizontalLegend />
      <CardGrid search>
        {
          searchResults.map((room, index) => {
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

export default SearchRooms
