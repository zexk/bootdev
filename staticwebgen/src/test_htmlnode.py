import unittest

from htmlnode import HTMLNode 

class TestHTMLNode(unittest.TestCase):
    def test_to_html_props(self):
        node = HTMLNode(
            "div",
            "Hello, world!",
            None,
            {"class": "greeting", "href": "https://google.com"},
        )
        self.assertEqual(node.props_to_html(),
            ' class="greeting" href="https://google.com"',
        )

    def test_values(self):
        node = HTMLNode(
            "div",
            "Test div",
        )
        self.assertEqual(
            node.tag,
            "div",
        )
        self.assertEqual(
            node.value,
            "Test div",
        )
        self.assertEqual(
            node.children,
            None,
        )
        self.assertEqual(
            node.props,
            None,
        )

    def test_repr(self):
        node = HTMLNode(
            "p",
            "Test p",
            None,
            {"class": "primary"},
        )
        self.assertEqual(
            node.__repr__(),
            "HTMLNode(p, Test p, None, {'class': 'primary'})",
        )

if __name__ == "__main__":
    unittest.main()
