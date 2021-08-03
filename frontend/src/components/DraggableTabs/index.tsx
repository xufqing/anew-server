import React from 'react';
import 'antd/dist/antd.css';
//import './index.css';
import { Tabs } from 'antd';
import { TabsProps } from 'antd/lib/tabs';
import { DndProvider, DragSource, DropTarget } from 'react-dnd';
import { HTML5Backend } from 'react-dnd-html5-backend';

function TabNode({ connectDragSource, connectDropTarget, children }: any) {
  return connectDragSource(connectDropTarget(children));
}
const cardTarget = {
  drop(props: any, monitor: any) {
    const dragKey = monitor.getItem().index;
    const hoverKey = props.index;

    if (dragKey === hoverKey) {
      return;
    }

    props.moveTabNode(dragKey, hoverKey);
    monitor.getItem().index = hoverKey;
  },
};

const cardSource = {
  beginDrag(props: any) {
    return {
      id: props.id,
      index: props.index,
    };
  },
};

const WrapTabNode = DropTarget('DND_NODE', cardTarget, (connect) => ({
  connectDropTarget: connect.dropTarget(),
}))(
  DragSource('DND_NODE', cardSource, (connect, monitor) => ({
    connectDragSource: connect.dragSource(),
    isDragging: monitor.isDragging(),
  }))(TabNode)
);


interface State {
  order: number[];
}
export class DraggableTabs extends React.Component<TabsProps, State> {
  constructor(props: TabsProps) {
    super(props);
    this.state = {
      order: [],
    };
  }

  moveTabNode = (dragKey: number, hoverKey: number) => {
    const tempState = this.state;
    const newOrder: number[] = tempState.order.slice();
    const { children } = this.props;

    React.Children.forEach(children, (c: any) => {
      if (newOrder.indexOf(c.key) === -1) {
        newOrder.push(c.key);
      }
    });

    const dragIndex = newOrder.indexOf(dragKey);
    const hoverIndex = newOrder.indexOf(hoverKey);

    newOrder.splice(dragIndex, 1);
    newOrder.splice(hoverIndex, 0, dragKey);

    this.setState({
      order: newOrder,
    });
  };

  renderTabBar = (props: any, DefaultTabBar: any) => (
    <DefaultTabBar {...props}>
      {(node: any) => (
        <WrapTabNode
          key={node.key}
          index={node.key}
          moveTabNode={this.moveTabNode}
        >
          {node}
        </WrapTabNode>
      )}
    </DefaultTabBar>
  );

  render() {
    const { order } = this.state;
    const { children } = this.props;

    const tabs: any[] = [];
    React.Children.forEach(children, (c) => {
      tabs.push(c);
    });

    const orderTabs = tabs.slice().sort((a, b) => {
      const orderA = order.indexOf(a.key);
      const orderB = order.indexOf(b.key);

      if (orderA !== -1 && orderB !== -1) {
        return orderA - orderB;
      }
      if (orderA !== -1) {
        return -1;
      }
      if (orderB !== -1) {
        return 1;
      }

      const ia = tabs.indexOf(a);
      const ib = tabs.indexOf(b);

      return ia - ib;
    });
    return (
      <DndProvider backend={HTML5Backend}>
        <Tabs renderTabBar={this.renderTabBar} {...this.props}>
          {orderTabs}
        </Tabs>
      </DndProvider>
    );
  }
}

export default DraggableTabs;