import React, {useState, useEffect} from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import Typography from '@mui/material/Typography';
import "./roomCard.css";

/* Props: 
 - RoomName
 - RoomHouse
 - Avaiability (Available, Booked, Booked)
 */
export default function RoomCardAlt(props) {

  let green = "#8ED264";
  let gray = "#808080";
  let red = "#E5414B";
  const [roomStatus, setRoomStatus] = useState(props.Avaiability);
  const [roomStatusColor, setRoomStatusColor] = useState(
    props.Avaiability === "available"
      ? green
      : props.Avaiability === "unknown"
      ? gray
      : red
      //Lägg till unknown
  );

  useEffect(() => {
    setRoomStatusColor(
      roomStatus === "available"
        ? green
        : roomStatus === "unknown"
        ? gray
        : red
    );
  }, [roomStatus])

  return (
    <Card sx={{ minWidth: 200, maxHeight: 160, maxWidth: 200 }} style={{borderRadius: "8px", margin: "8px"}}>
      <CardContent>
        <Typography variant="h5" component="div">
          {props.RoomName}
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          {props.RoomHouse}
        </Typography>
        <div 
        style={{
          width:"50%",
          marginLeft:"25%",
          marginTop:"15%",
          display:"flex",
          flexDirection:"row",
        }}>
          <Typography variant="body2" color="text.secondary">
            {roomStatus}
          </Typography>
        <div 
          className="mobile-room-card-availability-circle"
          style={{
            backgroundColor: roomStatusColor
          }} 
        />
      </div>
      </CardContent>
    </Card>
  );
}
