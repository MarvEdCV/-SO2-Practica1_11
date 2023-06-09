import React from 'react';
import './Homepage.css';
import { Container, Row, Col, Tab, Nav } from 'react-bootstrap';
import Statistics from '../Statistics/Statistics';
import AllProcess from '../AllProcess/AllProcess';
import Arbol from '../Arbol/Arbol';
import axios from 'axios';
export default class Homepage extends React.Component {
  constructor(props) {
    super(props);
    this.state = { Procesos: [], Arbol: [] };
  }
  componentDidMount() {
    this.interval = setInterval(() => {
      axios.get(this.props.URL + `/leercpu`)
        .then(res => {
          //console.log(res.data.procesos)
          this.setState({ Procesos: res.data.procesos });

          let arbol = res.data.arbol;
          arbol.map((proc) => {
            proc = this.NoCapitalLetter(proc)
            return null;
          })
          //console.log(arbol)
          this.setState({ Arbol: arbol });
        })
    }, 2000)
  }

  componentWillUnmount() {
    clearInterval(this.interval);
  }



  NoCapitalLetter(proc) {
    let obj = proc;
    obj.id = parseInt(obj.id);
    obj.label = "(" + obj.id + ") " + obj.label;
    obj.parentId = parseInt(obj.parentId);
    obj.items = obj.items;


    if (obj.parentId === 0) {
      obj.parentId = null
    } else if (obj.parentId === "null") {
      obj.parentId = null
    }

    if (obj.id == 899) {
      console.log(obj)
    }
    obj.items.map((hijo) => {
      hijo = this.NoCapitalLetter(hijo)
      return true;
    });


    return obj
  }

  render() {
    return (
      <Container fluid>
        <Statistics Procesos={this.state.Procesos} URL={this.props.URL} />
       <hr></hr>  
    
            <Tab.Container id="Selector-Tabs" defaultActiveKey="first" >
              
                  <Nav variant="pills" >
                    <Nav.Item className="pb-2">
                      <Nav.Link eventKey="first" className="btn btn-outline-danger"> PROCESS ALL</Nav.Link>
                    </Nav.Item>

                    <Nav.Item >
                      <Nav.Link eventKey="second" className="btn btn-outline-danger">PROCESS TREE</Nav.Link>
                    </Nav.Item>
                  </Nav>
               
                  <Tab.Content>
                    <Tab.Pane eventKey="first">
                      <AllProcess Procesos={this.state.Procesos} URL={this.props.URL} />
                    </Tab.Pane>
                    <Tab.Pane eventKey="second">
                      <Arbol Arbol={this.state.Arbol} URL={this.props.URL} />
                    </Tab.Pane>
                  </Tab.Content>
                
            </Tab.Container>
        

        





      </Container >
    );
  }
}

