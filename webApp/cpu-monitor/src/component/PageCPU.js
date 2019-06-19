import React, { Component } from 'react';
import ViewCPU from './sub/ViewCpu';
import ViewCPUHistory from './sub/ViewCPUHistory';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import { Redirect } from 'react-router-dom';

class PageCpu extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      redirectHome: false
    }
  }
  
  onSwapPage() {
    this.setState({
      redirectHome: true})
  }

  render() {
    if (this.state.redirectHome) {
      return <Redirect to="/"/>;
    } 
    else{
      return (
        <Container fluid={true}>
        <Row style={{ backgroundColor: 'black', height: '80px' }}>
          <Col lg={1}  className="text-left" >
            <Button variant="outline-light" style={{marginTop:'15px', height: '50px'}} onClick={() => {this.onSwapPage("/cpu")}}>Back</Button>
          </Col>
          <Col lg={11}>
          <p style={{margin:'0px', fontSize:'50px', color:'white'}}>CPU-State</p>
          </Col>
        </Row>
        <Row>
          <Col lg={5}><ViewCPU/></Col>
          <Col lg={7}>HISTORY</Col>
        </Row>
      </Container>
      );
    }
    
  }
}

export default PageCpu; // Don’t forget to use export default!