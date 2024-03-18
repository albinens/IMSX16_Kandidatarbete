import React, { useEffect, useState } from "react";

import "./roomCard.css";

/*
Props: 
 - RoomName
 - RoomHouse
 - Avaiability (Available, Booked, Occupied)
*/
function RoomCard(props) {
  let green = "#8ED264";
  let yellow = "#F4EC32";
  let red = "#E5414B";
  const [roomStatus, setRoomStatus] = useState(props.Avaiability);
  const [roomStatusColor, setRoomStatusColor] = useState(
    props.Avaiability === "available"
      ? green
      : props.Avaiability === "booked"
      ? yellow
      : red
  );

  useEffect(() => {
    setRoomStatusColor(
      roomStatus === "available"
        ? green
        : roomStatus === "booked"
        ? yellow
        : red
    );
  }, [roomStatus])

  return (
    <div className="mobile-room-card-root">
      <div className="mobile-room-card-info-text">
        <h3 className="mobile-room-card-name">{props.RoomName}</h3>
        <p className="mobile-room-card-house">{props.RoomHouse}</p>
      </div>
      <div 
        style={{
          width:"50%",
          marginLeft:"25%",
          marginTop:"10%",
        }}>
        <div 
          className="mobile-room-card-availability-circle"
          style={{
            backgroundColor: roomStatusColor
          }} 
        />
      </div>
    </div>
  );
}

export default RoomCard;