import { Variants } from 'framer-motion';

export const variants: Variants = {
  hiddenLeft: {
    opacity: 0,
    rotateY: '-45deg',
  },
  visible: {
    opacity: 1,
    rotateY: '0deg',
    transition: {
      type: 'tween',
      ease: 'easeOut',
      duration: 0.3,
    },
  },
  hiddenRight: {
    opacity: 0,
    rotateY: '45deg',
    transition: {
      type: 'tween',
      ease: 'easeIn',
      duration: 0.3,
    },
  },
};
