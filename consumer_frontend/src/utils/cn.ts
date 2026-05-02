/**
 * Utility function to combine className strings conditionally
 * Similar to 'classnames' library but minimal
 */
export function cn(...classes: (string | undefined | null | false)[]): string {
    return classes.filter(Boolean).join(' ');
}
