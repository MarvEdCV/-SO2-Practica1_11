import React from 'react';
import { Line } from 'react-chartjs-2';
import { Container, Row, Col, Table } from 'react-bootstrap';
import axios from 'axios';


export default class HOME extends React.Component {


    render() {
        return (
            <div>
                <h1>GRUPO 11 - PRACTICA 1 SISTEMAS OPERATIVOS 2</h1>
                <hr></hr>
                <Table bg="primary" hover striped>
                <thead bg="primary" >
                    <tr>
                        <th>Carnet</th>
                        <th>Nombre</th>
                        <th> </th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>201908053</td>
                        <td>Sara Paulina Medrano Cojulun</td>

                    </tr>
                    <tr>
                        <td>201905554</td>
                        <td>Marvin Eduardo Catalán Véliz</td>

                    </tr>

                    <tr>
                        <td>201709110</td>
                        <td>Wilson Eduardo Perez Echeverria</td>

                    </tr>
                </tbody>
            </Table>

            </div>
            
        );
    }
}
