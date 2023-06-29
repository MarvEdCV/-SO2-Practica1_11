import React, { useEffect, useState } from 'react';

import { Table, Button } from 'react-bootstrap';
import axios from 'axios';
import { useParams } from 'react-router-dom';


import Form from "react-bootstrap/Form";
export default function RamInfo() {
    useEffect(() => {
        // Lógica que deseas ejecutar al cargar la página
        prueba();
    }, []);
    const url = "https://918c-2803-d100-98a8-341e-ac5f-cfc6-c56b-a689.ngrok.io";
    const [maps, setMaps] = useState([]);
    const { id, nombre } = useParams(); // Obtener el ID de la URL
    const prueba = () => {
        axios.post(url + `/leermaps`, id)
          .then(res => {
            setMaps(res.data.maps);
          })
          .catch(error => {
            console.log(error);
          });
      };

      function convertirCadena(cadena) {
        const arreglo = Array.from(cadena); // Convertir la cadena en un arreglo de caracteres
        let resultado = "";

        arreglo.forEach((posicion, indice) => {
          // Realizar verificación condicional para cada posición del arreglo
          if (posicion === 'r') {
            resultado += "Lectura, ";
          } else if (posicion === 'w') {
            resultado += "Escritura, ";
          } else if (posicion === 'x') {
            resultado += "Ejecución, ";
          } else if (posicion === 'p') {
            resultado += "Privado/compartido ";
          } else {
            resultado += "";
          }
        });
        return resultado;
      }


    return (
        <div>
            <h2>{id} -- {nombre} -- Asignación de la memoria</h2>
            <Table bg="primary" hover striped>
                <thead bg="primary" >
                    <tr>
                        <th>Direcciones</th>
                        <th>Tamaños (KB)</th>
                        <th>Permisos</th>
                        <th>Dispositivos</th>

                        <th>Archivo </th>
                    </tr>
                </thead>
                <tbody>
                    {maps.map((map, key) => (
                        <tr key={key}>
                            <td>{map.addressRange}</td>
                            <td>{map.Size}</td>
                            <td>{convertirCadena(map.permissions)}</td>
                            <td>{map.device}</td>
                            <td>{map.path}</td>
                        </tr>
                    ))}

                </tbody>




            </Table>
        </div>


    );
}


