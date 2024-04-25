import React from "react";
import * as d3 from "d3";

const Table = ({ data }) => {

  d3.select('body').append('table')
  .style("border-collapse", "collapse")
  .style("border", "2px black solid");
  
  return (
    <>
    
    </>
  )
}

export default Table;