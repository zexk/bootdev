import unittest

from textnode import TextNode, TextType

class TestTextNode(unittest.TestCase):
    def test_eq(self):
        node = TextNode("This is a text node", TextType.BOLD)
        node2 = TextNode("This is a text node", TextType.BOLD)
        self.assertEqual(node, node2)

    def test_neq(self):
        node = TextNode("This is a text node", TextType.PLAIN)
        node2 = TextNode("This is a text node", TextType.BOLD)
        self.assertNotEqual(node, node2)

    def test_neq2(self):
        node = TextNode("This is a text node", TextType.PLAIN)
        node2 = TextNode("This is a text node 2", TextType.PLAIN)
        self.assertNotEqual(node, node2)

    def test_eq_url(self):
        node = TextNode("This is a text node", TextType.PLAIN, "https://www.google.com")
        node2 = TextNode("This is a text node", TextType.PLAIN, "https://www.google.com")
        self.assertEqual(node, node2)

    def test_repr(self):
        node = TextNode("This is a text node", TextType.PLAIN, "https://www.google.com")
        self.assertEqual(repr(node), "TextNode(This is a text node, TextType.PLAIN, https://www.google.com)")

if __name__ == "__main__":
    unittest.main()
