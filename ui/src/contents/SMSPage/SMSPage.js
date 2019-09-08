import React from 'react';
import { DataTableSkeleton } from 'carbon-components-react';
import SMSTable from './SMSTable';

const headers = [
  {
    key: 'cmgl_id',
    header: 'ID',
  },
  {
    key: 'mobile',
    header: 'Mobile',
  },
  {
    key: 'body',
    header: 'Text',
  },
  {
    key: 'type',
    header: 'Type',
  },
  {
    key: 'created_at',
    header: 'Created AT',
  }
];


class SMSPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      smss: [],
      loading: false,
      baseURI: "",
    };
    this.usageInit = this.usageInit.bind(this);
  }
  componentDidMount() {
    this.usageInit();
  }

  usageInit() {

    this.setState({ loading: true });
    const uri = this.state.baseURI + "/v1/smss";
    fetch(uri)
      .then(res => res.json())
      .then(result => {
        this.setState({ loading: false });
        if (result.smss != null) {
          this.setState({ smss: result.smss });
        }
      })
  }

  render() {
    if (this.state.loading) return (
      <DataTableSkeleton
        columnCount={headers.length + 1}
        rowCount={10}
        headers={headers}
      />);

    return (
      <SMSTable headers={headers} rows={this.state.smss} />
    );
  };
}

export default SMSPage;