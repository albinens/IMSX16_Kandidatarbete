import React from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import './App.css'
import ListRooms from './pages/listRooms'
import SearchRooms from './pages/searchRooms'
import Layout from './layout';



function App() {
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<ListRooms />}/>
          <Route path="search" element={<SearchRooms />}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
