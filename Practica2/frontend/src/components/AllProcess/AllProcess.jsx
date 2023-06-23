import React from 'react';
import './AllProcess.css';
import { Table, Button } from 'react-bootstrap';
import axios from 'axios';
import { Link } from 'react-router-dom';

import Form from "react-bootstrap/Form";
export default class AllProcess extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      filtroNombre: "" // Estado para almacenar el valor del filtro por Nombre
    };

    //this.killprocess = this.killprocess.bind(this);
  }
  handleChange = (event) => {
    this.setState({ filtroNombre: event.target.value });
  };
  killprocess(e) {
    axios.post(this.props.URL + `/killprocess`, e.target.value)
      .then(res => {
        console.log(res.data);
      })
  }
  prueba(e) {
    axios.post(this.props.URL + `/leermaps`, e.target.value)
      .then(res => {
        console.log(res.data);
      })
  }
  render() {
    const { filtroNombre } = this.state;
    const procesosFiltrados = this.props.Procesos.filter(
      (process) =>
        process.Nombre.toLowerCase().includes(filtroNombre.toLowerCase()) // Filtrar por coincidencia de nombre
    );
    return (
      <div>
        <Form.Group controlId="filtroNombre">
          <Form.Label>Buscar por Nombre:</Form.Label>
          <Form.Control
            type="text"
            value={filtroNombre}
            onChange={this.handleChange}
          />
        </Form.Group>
        <Table bg="primary" hover striped>
        <thead bg="primary" >
          <tr>
            <th>PID</th>
            <th>Usuario</th>
            <th>Nombre</th>
            <th>Estado</th>
           
            <th> </th>
          </tr>
        </thead>
        {procesosFiltrados.map((process, key) => {
              return (
                <tr key={key}>
                  <td>{process.PID}</td>
                  <td>{process.Usuario}</td>
                  <td>{process.Nombre}</td>
                  <td>{process.Estado === 1 ? 'Ejecucion' : process.Estado === 4 ? 'Parado' : process.Estado === 1? 'Durmiendo': process.Estado === 1026? 'Durmiendo': 'Zombie'}</td>
                  
                  <td>
                    <Button
                      id={process.PID}
                      value={process.PID}
                      variant="danger"
                      onClick={this.killprocess.bind(this)}
                    >
                      kill
                    </Button>
                    
                  </td>
                  <td>
                 
                  <Button
                      id={process.PID}
                      value={process.PID}
                      variant="warning"
                     
                    >
                      <Link to={`/INFORAM/${process.PID}/${encodeURIComponent(process.Nombre)}`}>RAM</Link>
                    </Button>

                  </td>
                </tr>
              );
            })}
      </Table>
      </div>
     

    );
  }
}

