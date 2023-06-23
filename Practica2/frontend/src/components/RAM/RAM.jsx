import React from 'react';
import './RAM.css';
import { Line } from 'react-chartjs-2';
import { Container, Row, Col, Table } from 'react-bootstrap';
import axios from 'axios';


export default class RAM extends React.Component {
  constructor(props) {
    super(props);

    let data_percentage = {
      labels: [],
      datasets: [
        {
          label: 'Total de RAM ',
          fill: false,
          lineTension: 0.1,
          backgroundColor: 'rgba(75,192,192,0.4)',
          borderColor: 'rgba(75,192,192,1)',
          borderCapStyle: 'butt',
          borderDash: [],
          borderDashOffset: 0.0,
          borderJoinStyle: 'miter',
          pointBorderColor: 'rgba(75,192,192,1)',
          pointBackgroundColor: 'rgba(75,192,192,1)',
          pointBorderWidth: 1,
          pointHoverRadius: 5,
          pointHoverBackgroundColor: 'rgba(75,192,192,1)',
          pointHoverBorderColor: 'rgba(220,220,220,1)',
          pointHoverBorderWidth: 2,
          pointRadius: 5,
          pointHitRadius: 10,
          data: []
        }
      ]
    };


    this.state = { labels: [], data: data_percentage.datasets[0].data, actual: 0, };
  }


  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(this.props.URL + `/leerram`)
        .then(res => {
          console.log(res.data)

          this.setState({
            RAMConsumida: res.data.MemoriaTotal - res.data.MemoriaLibre,
            percentage_ram: res.data.MemoriaUsada,
            TotalRam: res.data.MemoriaTotal
          });

          let resp = res.data.MemoriaTotal - res.data.MemoriaLibre;

          let labels1 = this.state.labels;
          let dt = new Date();
          labels1.push(dt.toLocaleTimeString());
          let data1 = this.state.data;

          data1.push(resp);
          if (labels1.length > 10) {
            labels1.shift();
            data1.shift();
          }

          this.setState({ data: data1 })
          //dataset
          let data_percentage = {
            labels: labels1,
            datasets: [
              {
                label: 'Total de RAM (MB)',
                fill: false,
                lineTension: 0.1,
                backgroundColor: 'rgba(220, 53, 69, 0.4)',
                borderColor: 'rgba(220, 53, 69, 0.4)',
                borderCapStyle: 'butt',
                borderDash: [],
                borderDashOffset: 0.0,
                borderJoinStyle: 'miter',
                pointBorderColor:'rgba(220, 53, 69, 0.4)',
                pointBackgroundColor: 'rgba(220, 53, 69, 0.4)',
                pointBorderWidth: 1,
                pointHoverRadius: 5,
                pointHoverBackgroundColor: 'rgba(220, 53, 69, 0.4)',
                pointHoverBorderColor: 'rgba(220, 53, 69, 0.4)',
                pointHoverBorderWidth: 2,
                pointRadius: 5,
                pointHitRadius: 10,
                data: data1
              }
            ]
          };

          this.setState({ actual: resp, labels: labels1, data: data1, dataset: data_percentage });
          let lineChart = this.reference.chartInstance
          lineChart.update();
        })
    }, 3000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    return (
      <div>
        <Container fluid>
          <Row>
            <Col md={5}>
            </Col>
            <Col md={7}>
              <h1 > CONTROL DE RAM</h1>
            </Col>
          </Row>
         
            <Col md={8}>
              <Line data={this.state.dataset} ref={(reference) => this.reference = reference} />
            </Col>
            <h1>DATOS EN TIEMPO REAL</h1>
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
              <td> Cantidad Total de la RAM</td>
              <td> {this.state.TotalRam}MB</td>

            </tr>
            <tr>
              <td> Cantidad consumida de la RAM</td>
              <td> {this.state.RAMConsumida}MB</td>

            </tr>

            <tr>
              <td>  Porcentaje  de Consumo de la RAM</td>
              <td> {this.state.percentage_ram}%</td>

            </tr>
           
          </tbody>
        </Table>

        </Container>

      </div>
    );
  }
}
