import { getNameInitials } from "@/utils/get-name-initials";
import { Avatar as AntdAvatar } from "antd";

type Props = AvatarProps & {
  name?: string;
};

const CustomAvatar = ({ name, style, ...rest }: Props) => {
  return (
    <AntdAvatar
      alt={"Astrobear"}
      size="small"
      style={{
        backgroundColor: "#88d068",
        display: "flex",
        alignItems: "center",
        border: "none",
        ...style,
      }}
      {...rest}
    >
      {getNameInitials(name || "")}
    </AntdAvatar>
  );
};

export default CustomAvatar;
