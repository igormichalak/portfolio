import styled from 'styled-components';
import { motion } from 'framer-motion';

export interface WrapperProps {
  $topMargin?: boolean;
  $centerContent?: boolean;
}

export const Wrapper = styled(motion.div)<WrapperProps>`
  width: 100%;
  height: 100%;
  transform-origin: center center -927px;
  transform-style: preserve-3d;
  display: flex;
  flex-direction: column;

  ${({ $topMargin }) => ($topMargin ? `margin-top: 128px;` : ``)}

  ${({ $centerContent }) => ($centerContent ? `align-items: center;` : ``)}
`;
