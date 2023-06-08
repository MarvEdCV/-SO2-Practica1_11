import React from 'react';
import './Arbol.css';
import axios from 'axios';
import Tree from '@naisutech/react-tree'
export default class Arbol extends React.Component {
  constructor(props) {
    super(props);

    this.state = { Procesos: [] };

  }
  

  render() {
    
    return (

      <div>
        <Tree
          nodes={this.props.Arbol} 
          className="custom-tree"
          id="tree"
        />
      </div>
    );
  }
}

