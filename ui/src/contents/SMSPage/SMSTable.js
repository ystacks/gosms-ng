import React from 'react';
import {
    DataTable,
    TableContainer,
    TableToolbar,
    TableToolbarContent,
    TableBatchActions,
    TableBatchAction,
    Table,
    TableHead,
    TableRow,
    TableExpandHeader,
    TableHeader,
    TableBody,
    TableExpandRow,
    TableCell,
    TableExpandedRow,
    Button,
    TableSelectAll,
    TableSelectRow,
} from 'carbon-components-react';

function smsDelete(selectedRows) {
    const uri = "/v1/devices/1/sms";
    var smsids = [];
    var payload = {};

    selectedRows.forEach(function(row) {
        smsids.push(row["id"]);
      });

    payload["sms_ids"] = smsids;
    payload["action"] = "delete";

    fetch(uri, {
        method: 'PATCH',
        mode: 'cors',
        cache: 'no-cache',
        headers: {
          'Content-Type': 'application/json',
        },
        redirect: 'follow',
        referrer: 'no-referrer',
        body: JSON.stringify(payload),
      })
        .then(response => response.json())
        .then(data => {
            console.log(data);
        })
        .catch(error => console.error(error));

}

const SMSTable = ({ rows, headers }) => {
    return (
        <DataTable
            rows={rows}
            headers={headers}
            render={({
                rows,
                headers,
                getHeaderProps,
                getRowProps,
                getTableProps,
                getSelectionProps,
                getBatchActionProps,
                overflowMenuProps,
                overflowMenuItemProps,
                selectedRows,
            }) => (
                    <TableContainer >
                        <TableToolbar>
                            <TableBatchActions {...getBatchActionProps()}>
                                <TableBatchAction onClick={() => smsDelete(selectedRows)}>
                                    Delete
                                </TableBatchAction>
                            </TableBatchActions>
                            <TableToolbarContent>
                                {(
                                    <Button size="small" kind="primary" >Compose</Button>
                                )}
                            </TableToolbarContent>
                        </TableToolbar>
                        <Table {...getTableProps()}>
                            <TableHead>
                                <TableRow>
                                    <TableExpandHeader />
                                    <TableSelectAll {...getSelectionProps({})} />
                                    {headers.map(header => (
                                        <TableHeader {...getHeaderProps({ header })}>
                                            {header.header}
                                        </TableHeader>
                                    ))}
                                </TableRow>
                            </TableHead>
                            <TableBody>
                                {rows.map(row => (
                                    <React.Fragment key={row.id}>
                                        <TableExpandRow {...getRowProps({ row })}>
                                            <TableSelectRow {...getSelectionProps({ row })} />
                                            {row.cells.map(cell => (
                                                <TableCell key={cell.id}>{cell.value}</TableCell>
                                            ))}
                                        </TableExpandRow>
                                        <TableExpandedRow colSpan={headers.length + 2}>
                                            <p>Row description</p>
                                        </TableExpandedRow>
                                    </React.Fragment>
                                ))}
                            </TableBody>
                        </Table>
                    </TableContainer>
                )}
        />
    );
};
export default SMSTable;