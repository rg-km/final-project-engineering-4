import { Box } from "@chakra-ui/react";
import { Link } from "react-router-dom";

const MenuItem = ({label, redirect}) => {
    return <Box as={Link} to={redirect} minW={"200px"}>
        <Box
            bg={'white'}
            py={{ base: 1, md: 2 }}
            px={{ base: 2, md: 4 }}
            rounded={{ base: "xl", md: "2xl" }}
            shadow={'sm'}
            border={'1px'}
            borderColor={'transparent'}
            textAlign={'center'}
            fontSize={{base: "xs", md: "md"}}>
            {label}
        </Box>
    </Box>
}

export default MenuItem;