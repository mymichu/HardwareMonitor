import React, { Component } from 'react';
import Container from 'react-bootstrap/Container';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import { Redirect } from 'react-router-dom';

const buttonStyle = {
  width:'250px', 
  height:'250px', 
  margin: 'auto',
};


class PageHome extends Component {
  constructor(props) {
    super(props);
    this.state = { 
      url: "",
      redirect: false
    }
  }
  
  onSwapPage(path) {
    this.setState({url:path,
    redirect: true})
  }

  render() {
    if (this.state.redirect) {
      return <Redirect to={this.state.url}/>;
    } 
    return (
      <Container fluid={true}>
        <Row style={{ backgroundColor: 'black', height: '80px' }}>
          <Col lg={12}>
          <p style={{margin:'0px', fontSize:'50px', color:'white'}}>Demo</p>
          </Col>
        </Row>
        <Row style={{padding:"10px"}}>
          <Col lg={4} className="text-center">
            <Button variant="outline-dark" style={buttonStyle} onClick={() => {this.onSwapPage("/cpu")}}>CPU</Button>
          </Col>
          <Col lg={4} className="text-center">
            <Button variant="outline-dark" style={buttonStyle}>Temp</Button>
            </Col>
            <Col lg={4} className="text-center">
            <Button variant="outline-dark" style={buttonStyle} disabled>-</Button>
            </Col>
        </Row>
      </Container>
    );
  }
}

export default PageHome; // Don’t forget to use export default!