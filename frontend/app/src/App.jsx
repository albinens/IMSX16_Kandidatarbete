import React from 'react'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import './App.css'
import ListRooms from './pages/listRooms'
import SearchRooms from './pages/searchRooms'
import Layout from './layout';
import Settings from './pages/settings';
import Sensors from './pages/sensors';
import DataBoard from './pages/dataBoard';

function App() {
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<ListRooms />}/>
          <Route path="search" element={<SearchRooms />}/>
          <Route path="settings" element={<Settings />}/>
          <Route path='sensors' element={<Sensors />}/>
          <Route path='databoard' element={<DataBoard />}/>
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App
