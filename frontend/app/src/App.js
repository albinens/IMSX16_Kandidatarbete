import './App.css';
import CardGrid from './components/cardGrid/cardGrid';
import MobileRoomCard from './components/mobileRoomCard/mobileRoomCard';

function App() {
  return (
    <div className="App">
      <CardGrid>
        <MobileRoomCard 
          RoomName="Vasa G-14"
          RoomHouse="Vasa"
          Avaiability="Available"
        />
        <MobileRoomCard 
          RoomName="Vasa G-15"
          RoomHouse="Vasa"
          Avaiability="Booked"
        />
        <MobileRoomCard 
          RoomName="Vasa G-16"
          RoomHouse="Vasa"
          Avaiability="Occupied"
        />
      </CardGrid>
    </div>
  );
}

export default App;
