import React, { Component } from 'react';
import {
  CircularProgressbar,
  buildStyles
} from "react-circular-progressbar";
import "react-circular-progressbar/dist/styles.css";
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';

class ViewCpu extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      isApiFetched: false,
      temp: 0.0,
      usage: 0
     };
  }

  componentDidMount(){
    this.interval = setInterval(() => fetch('http://192.168.10.1:82/api/v1/cpu/state/current')
    .then(res => res.json())
    .then(json => {
      this.setState({
        isApiFetched: true,
        temp: json.Temp,
        usage: json.Usage,
      }
      )
    }), 1000);
  }
  componentWillUnmount() {
    clearInterval(this.interval);
  }

  render() {
    
    return (
         <Container fluid={true}>
         <Row style={{ height: '60px' }}>
         <Col lg={6}>
          <p style={{margin:'0px', fontSize:'30px', color:'grey'}}>Usage</p>
          </Col>
          <Col lg={6}>
          <p style={{margin:'0px', fontSize:'30px', color:'grey'}}>Temperature</p>
          </Col>
        </Row>
        <Row>
          <Col lg={6}><CircularProgressbar styles={buildStyles({
              pathTransitionDuration: 0.15
            })}
            value={this.state.usage} 
            text={`${this.state.usage}%`} />  
          </Col>
          <Col lg={6}>
            <CircularProgressbar styles={buildStyles({
              pathTransitionDuration: 0.15
            })}
            value={this.state.temp} 
            text={`${this.state.temp}°C`} 
            maxValue={100}
            />  

          </Col>
        </Row>
      </Container>
    );
  }
}

export default ViewCpu; // Don’t forget to use export default!