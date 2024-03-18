import React, {useEffect, useState} from 'react'
import axios from 'axios'
import CardGrid from '../components/cardGrid/cardGrid'
import RoomCard from '../components/roomCard/roomCard'
import HorizontalLegend from '../components/legends/horizontalLegend/horizontalLegend'
import './styles/search.css'

function SearchRooms() {

  const [searchQuery, setSearchQuery] = useState('')
  const [apiData, setApiData] = useState([]) // [roomName, house, avaiability
  const [searchResults, setSearchResults] = useState([])


  const client = axios.create({
    baseURL: "http://localhost:8080/api",
  })

  const handleSearch = (e) => {
      setSearchQuery(e.target.value)
      console.log(searchQuery)
      if(e.target.value === ''){
        setSearchResults(apiData) 
      }
  }

  const handleSearchButton = () => {
    let filteredData = apiData.filter((room) => { //Change to API data
      return room.room.toLowerCase().includes(searchQuery.toLowerCase())
    })
    setSearchResults(filteredData)
    console.log(searchResults)
  }

  const handleEnterSearch = (e) => {
    if (e.key === 'Enter') {
      handleSearchButton()
    }
  }

  useEffect(() => {
    // fetch data from api
    const fetchData = async () => {
      client.get('/current').then((response) => { 
        setApiData(response.data);
      });
    }
    fetchData()
    setSearchResults(apiData) 
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
            if (room.status !== "") {
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

export default SearchRooms
