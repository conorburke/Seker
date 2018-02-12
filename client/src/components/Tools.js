import React, { Component } from 'react';
import { connect } from 'react-redux';

class Tools extends Component {
  render() {
    console.log(this.props);
    console.log('tools', this.props.tools);
    const allTools = this.props.tools.map(tool => {
        console.log(tool);
        return <li key={tool.ID}>{tool.Name}</li>;
      }
    )

    return (
      <div>
        <h1>tools</h1>
        <ul>{allTools}</ul>
      </div>
      ) 
  }
}

function mapStateToProps(state) {
  return { tools: state.tools };
}

export default connect(mapStateToProps)(Tools);