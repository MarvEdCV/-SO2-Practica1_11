import React from 'react';
import './Statistics.css';
import { Container, Row, Col, Tab, Nav, Badge, Table } from 'react-bootstrap';
import axios from 'axios';
export default class Statistics extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Total: 0, Running: 0, Suspend: 0, Stop: 0, Zombie: 0 };
    /*axios.get(this.props.URL + `/statistics`)
      .then(res => {
        this.setState({ Total: res.data.Total, Running: res.data.Running, Suspend: res.data.Suspend, Stop: res.data.Stop, Zombie: res.data.Zombie });
      })*/

  }
  componentDidMount() {
    this.interval = setInterval(() => {
      /*axios.get(this.props.URL + `/statistics`)
        .then(res => {
          this.setState({ Total: res.data.Total, Running: res.data.Running, Suspend: res.data.Suspend, Stop: res.data.Stop, Zombie: res.data.Zombie });
        })*/
      console.log(this.props.Procesos)
      let running = this.props.Procesos.filter((proces) => proces.Estado == 0 && proces.Memory != null).length;
      let stop = this.props.Procesos.filter((proces) => proces.Estado == 4).length;
      let zombie = this.props.Procesos.filter((proces) => proces.Estado == 4).length;
      let suspend = this.props.Procesos.filter((proces) => (proces.Estado == 1 || proces.Estado == 1026) && proces.Memory != null).length;
      this.setState({
        Total: this.props.Procesos.filter((proces) => proces.Memory != null).length,
        Running: running,
        Stop: stop,
        Suspend: suspend,
        Zombie: zombie
      });
    }, 1000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>

        <h1>CANTIDAD DE PROCESOS SEGUN DESCRIPCIÓN</h1>
       
        <Table bg="primary" hover striped>
          <thead bg="primary" >
            <tr>
              <th>DESCRIPCIÓN</th>
              <th>CANTIDAD</th>
              <th> </th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Total de Procesos</td>
              <td>{this.state.Total}</td>

            </tr>
            <tr>
              <td> Procesos en Ejecucion</td>
              <td>{this.state.Running}</td>

            </tr>

            <tr>
              <td> Procesos Durmiendo</td>
              <td>{this.state.Suspend}</td>

            </tr>
            <tr>
              <td> Procesos Parados</td>
              <td>{this.state.Stop}</td>

            </tr>
            <tr>
              <td> Procesos Zombie</td>
              <td>{this.state.Zombie}</td>

            </tr>
          </tbody>
        </Table>

      </div>
    );
  }
}

