import styled from 'styled-components';
import { motion } from 'framer-motion';

export const Wrapper = styled.li`
  position: relative;
  display: flex;
  align-items: center;
`;

export const ActiveIndicator = styled(motion.div)`
  position: absolute;
  inset: 0 auto auto 0;
  width: 4px;
  height: 24px;
  background-color: ${({ theme }) => theme.colors.primary[3]};
  border-radius: 2px;
`;

export const Link = styled.a<{ isActive: boolean }>`
  margin-left: 24px;
  font-size: 16px;
  font-weight: 600;
  line-height: 1.5;

  ${({ isActive, theme }) =>
    !isActive ? `color: ${theme.colors.neutral[4]};` : ``}

  &:hover {
    cursor: pointer;
  }
`;
