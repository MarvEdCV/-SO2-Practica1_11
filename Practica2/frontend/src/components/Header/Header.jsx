import React from 'react';
import './Header.css';
import { Navbar, Nav } from 'react-bootstrap';


export default class Header extends React.Component {
  render() {
    return (
      <div>
        <Navbar bg="danger" variant="dark">
          <Navbar.Brand href="/">Control, Creación y Monitoreo de Procesos</Navbar.Brand>
          <Nav className="mr-auto">
            <Nav.Link href="/CPU">CPU</Nav.Link>
            <Nav.Link href="/RAM_Monitor">RAM</Nav.Link>
          </Nav>
        </Navbar>
      </div>
    );
  }
}
