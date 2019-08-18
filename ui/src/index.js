/**
 * @file             : index.js
 * @author           : Jiang Yitao <jiangyt.cn#gmail.com>
 * Date              : 18.08.2019
 * Last Modified Date: 18.08.2019
 * Last Modified By  : Jiang Yitao <jiangyt.cn#gmail.com>
 */
import ReactDOM from 'react-dom';
import Spinner from 'react-bootstrap/Spinner';

ReactDOM.render(
  <div>
  <Spinner animation="border" variant="primary" />
  <Spinner animation="border" variant="secondary" />
  <Spinner animation="border" variant="success" />
  <Spinner animation="border" variant="danger" />
  <Spinner animation="border" variant="warning" />
  <Spinner animation="border" variant="info" />
  <Spinner animation="border" variant="light" />
  <Spinner animation="border" variant="dark" />
  </div>,
  document.getElementById('content')
);
