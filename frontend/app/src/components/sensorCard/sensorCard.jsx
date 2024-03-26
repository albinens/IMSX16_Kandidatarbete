import React, { useEffect, useState } from "react";

import "./sensorCard.css";

/*
Props: 
 - RoomName
 - RoomHouse
 - Avaiability (Available, Booked, Occupied)
*/
function SensorCard(props) {
  let green = "#8ED264";
  let yellow = "#F4EC32";
  let red = "#E5414B";
  const [sensorStatus, setsensorStatus] = useState(props.Avaiability);
  const [sensorStatusColor, setsensorStatusColor] = useState(
    props.Avaiability === "Online"
      ? green
      : props.Avaiability === "booked"
      ? yellow
      : red
  );

  useEffect(() => {
    setsensorStatusColor(
      sensorStatus === "Online"
        ? green
        : sensorStatus === "booked"
        ? yellow
        : red
    );
  }, [sensorStatus])

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
            backgroundColor: sensorStatusColor
          }} 
        />
      </div>
    </div>
  );
}

export default SensorCard;