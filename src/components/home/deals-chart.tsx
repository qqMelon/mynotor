import { Card } from "antd";
import { DollarOutlined } from "@ant-design/icons";
import { Text } from "../text";

const DealsChart = () => {
  return (
    <Card
      style={{ height: "100%" }}
      title={
          <div>
            <DollarOutlined />
              <Text size="sm" style={{ marginLeft: " 0.5rem" }}>
                Deals
              </Text>
          </div>
      }
    >
    </Card>
  )
};

export default DealsChart;
